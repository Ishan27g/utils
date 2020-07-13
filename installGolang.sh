#wget latest golang for linux - https://golang.org/dl/
# run with .tar.gz file as 1st argument

tar -xvf $1
sudo chown -R root:root ./go
sudo mv go /usr/local
Home=`echo $HOME`
echo "Run following commands"
echo ""
echo "export GOPATH=$Home/go"
echo ""
echo "export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin"
echo ""
