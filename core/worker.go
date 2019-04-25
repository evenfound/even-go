package core

import (
	"os"
	"os/signal"
	"path/filepath"

	l "github.com/evenfound/even-go/utility/log"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
)

var (
	// Init and setup logger.
	logger = l.NewBackend(os.Stdout).Logger("CORE")

	// shutdownRequestChannel is used to initiate shutdown from one of the subsystems using the same code paths
	// as when an interrupt signal is received.
	shutdownRequestChannel = make(chan struct{})

	// interruptSignals defines the default signals to catch in order to do a proper shutdown.
	// This may be modified during init depending on the platform.
	interruptSignals = []os.Signal{os.Interrupt}
)

// interruptListener listens for OS Signals such as SIGINT (Ctrl+C) and shutdown requests from shutdownRequestChannel.
// It returns a channel that is closed when either signal is received.
func interruptListener() <-chan struct{} {

	c := make(chan struct{})

	go func() {

		interruptChannel := make(chan os.Signal, 1)
		signal.Notify(interruptChannel, interruptSignals...)

		// Listen for initial shutdown signal and close the returned channel to notify the caller.
		select {
		case sig := <-interruptChannel:
			logger.Infof("Received signal (%s).", sig)

		case <-shutdownRequestChannel:
			logger.Info("Shutdown requested. Shutting down...")
		}

		close(c)

		// Listen for repeated signals and display a message so the user knows the shutdown is in progress and the process is not hung.
		for {
			select {
			case sig := <-interruptChannel:
				logger.Infof("Received signal (%s). Already shutting down.", sig)

			case <-shutdownRequestChannel:
				logger.Info("Shutdown requested. Already shutting down.")
			}
		}

	}()

	return c
}

// interruptRequested returns true when the channel returned by interruptListener was closed.
// This simplifies early shutdown slightly since the caller can just use an if statement instead of a select.
func interruptRequested(interrupted <-chan struct{}) bool {

	select {
	case <-interrupted:
		return true
	default:
	}

	return false
}

func Worker() error {
	// Get a channel that will be closed when a shutdown signal has been
	// triggered either from an OS signal such as SIGINT (Ctrl+C) or from
	// another subsystem such as the RPC server.
	interrupt := interruptListener()

	// Return now if an interrupt signal was triggered.
	if interruptRequested(interrupt) {
		return nil
	}

	defer func() {

		logger.Info("Gracefully shutting down the Even Network...")

		if Node != nil {

			if Node.MessageRetriever != nil {
				Node.RecordAgingNotifier.Stop()
				logger.Info("RecordAgingNotifier stopped...")
				close(Node.MessageRetriever.DoneChan)
				Node.MessageRetriever.Wait()
				logger.Info("MessageRetriever closed...")
			}

			OfflineMessageWaitGroup.Wait()
			Node.Datastore.Close()
			logger.Info("Data-store unlocked and closed...")
			os.Remove(filepath.Join(Node.RepoPath, fsrepo.LockFile))

			err := Node.IpfsNode.Close()
			if err != nil {
				logger.Critical(err)
			}
			logger.Info("IPFS-Node closed...")

			logger.Info("Even Network shutdown completed.")
		}

	}()

	// Wait until the interrupt signal is received from an OS signal or shutdown
	// is requested through one of the subsystems such as the RPC server.
	<-interrupt

	return nil
}
