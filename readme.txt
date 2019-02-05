1. Install MSYS2 (Mingw64);
2. Install Visual Studio Code;
3. Set Msys2 terminal for VSCode terminal default;
4. PLace code:

cd () {
    builtin cd "$@"
    cdir=$PWD
    while [ "$cdir" != "/" ]; do
        if [ -e "$cdir/.gopath" ]; then
            export GOPATH=$cdir
            break
        fi
        cdir=$(dirname "$cdir")
    done
    cdir=$PWD
    while [ "$cdir" != "/" ]; do
        if [ -e "$cdir/.gobin" ]; then
            export PATH=$PATH:$cdir
            break
        fi
        cdir=$(dirname "$cdir")
    done
}

to ~/.bash_profile end;

5. Launch terminal Msys2, type and run command: source ~/.bash_profile;
6. Create file named .gopath in the every path, that you want make it as GOPATH;
7. Clone projects in to choosen GOPATH directory after

go get -u github.com/evenfound/even-go 

in his root running;

8. For developers: change the git url in the .git/config [remote "origin"] as:
url = git@github.com:evenfound/even-go.git

9. Make fork ipfs refactored repository as:

mkdir $GOPATH/src/github.com/ipfs
git clone git@github.com:OpenBazaar/go-ipfs.git

10. Install godep:
go get github.com/tools/godep
 