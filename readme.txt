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

9. Install godep:
go get github.com/tools/godep


10. Install ipfs and gx (automatically). For windows (on 10 sample show) need run follow:
- Type "secpol.msc" in RUN dialog box and press Enter. It'll open Local Security Policy.
- Now go to "Local Policies -> Security Options". In right-side pane, scroll down to last and you'll see following options related to UAC:
- User Account Control: Admin Approval Mode for the Built-in Administrator account;
- Tap "Enable"; 
- Restart Windows;
After that run in Msys2 console

$ go get -u -d github.com/ipfs/go-ipfs
$ cd $GOPATH/src/github.com/ipfs/go-ipfs
$ make install

After first installation that update gx repo link make type 
gx install --global

11. Enter to even-go root and run command:
godep get
godep save 
 