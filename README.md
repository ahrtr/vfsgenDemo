vfsgenDemo
======
Firstly, you need to build the demo using the following commands,

```bash
git clone https://github.com/ahrtr/vfsgenDemo vfsgenDemo
cd vfsgenDemo
./build.sh
```

You will see that a new file named "assets_vfsdata.go" is generated under directory "data", which statically implements the virtual filesystem. The main.go is a demo on how to use the virtual filesystem. You can run the demo using the following command,

```bash
bin/vfsgen-demo
```
