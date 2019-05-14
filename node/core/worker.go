package core

import (
	"fmt"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	"os"
	"os/signal"
	"path/filepath"
)

// shutdownRequestChannel is used to initiate shutdown from one of the subsystems using the same code paths
// as when an interrupt signal is received.
var shutdownRequestChannel = make(chan struct{})

// interruptSignals defines the default signals to catch in order to do a proper shutdown.
// This may be modified during init depending on the platform.
var interruptSignals = []os.Signal{os.Interrupt}

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
			fmt.Printf("\n[INF] SGNL: Received signal (%s).\n", sig)

		case <-shutdownRequestChannel:
			fmt.Println("\n[INF] SGNL: Shutdown requested. Shutting down...")
		}

		close(c)

		// Listen for repeated signals and display a message so the user knows the shutdown is in progress and the process is not hung.
		for {
			select {
			case sig := <-interruptChannel:
				fmt.Printf("[INF] SGNL: Received signal (%s). Already shutting down.\n", sig)

			case <-shutdownRequestChannel:
				fmt.Println("[INF] SGNL: Shutdown requested. Already shutting down.")
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

		fmt.Println("[INF] EVNET: Gracefully shutting down the Even Network...")

		if Node != nil {
			if Node.MessageRetriever != nil {
				Node.RecordAgingNotifier.Stop()
				fmt.Println("[INF] EVNET: RecordAgingNotifier stopped...")
				close(Node.MessageRetriever.DoneChan)
				Node.MessageRetriever.Wait()
				fmt.Println("[INF] EVNET: MessageRetriever closed...")
			}

			OfflineMessageWaitGroup.Wait()
			Node.Datastore.Close()
			fmt.Println("[INF] EVNET: Data-store unlocked and closed...")
			os.Remove(filepath.Join(Node.RepoPath, fsrepo.LockFile))

			err := Node.IpfsNode.Close()
			if err != nil {
				log.Panic(err)
			}
			fmt.Println("[INF] EVNET: IPFS-Node closed...")

			fmt.Println("\n[EXIT] EVNET: Even Network shutdown completed.")
		}

	}()

	// Wait until the interrupt signal is received from an OS signal or shutdown
	// is requested through one of the subsystems such as the RPC server.
	<-interrupt

	return nil

}
