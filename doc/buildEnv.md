### Eclipse用惯了，所以开发环境采用eclipse+go插件的方式搭建。 参考：http://goclipse.github.io/

* 1.安装最新的go，配置goroot、gopath

* 2.下载eclipse，因为需要debug功能，而debug需要gdb，所以建议下载Eclipse IDE for C/C++ Developers

* 3.下载完毕后，如果配置了科学上网，那么可以直接 Help -> Install New Software -> add -> http://goclipse.github.io/releases/ 来进行插件的安装。如果没有配置科学上网，那么就先把go插件下载下来，下载地址为https://github.com/GoClipse/goclipse.github.io/archive/master.zip， 解压后，用releases那一层的目录作为 Local repository，接着安装即可，注意需要Uncheck the option “Contact all updates sites during installation to find required software”

* 4.设置环境变量Windows->Reference-> Go 这一步主要配置是配置gopath、goroot，在tools的选项里面，配置gocode、guru、godef，配置后以上后点击apply。对于gocode、guru、godef 这3个工具需要进行编译，编译的方法是下载下来后，执行go build xxxx，然后生成对应的 .exe 文件。在这儿提供一个编译好的下载 http://pan.baidu.com/s/1qYAvUFi

* 5.配置debug功能，就是想gdb配置好。网上说下载MinGW（http://sourceforge.net/projects/mingw/files/MinGW/ ）， 然后下载相应的gdb包，再在相应的gdb debugger写上gdb文件的路径。我试了试，一直在说版本不对，所以google了一下，最终的解决方案是下载一个liteide，在它的bin文件夹里面有gdb64.exe，将这个路径填入上述的地方即可。

* 6.将eclipse中的go project上传到github的步骤：
	* team shared project --> 勾选  Use or create repository in parent folder of project --> 点击  Create Repository  ->  Finish
	* Team --> commit 本地提交
	* 在github建立一个你想命名的github project，如果缺少这一步，有可能在下一步push的时候，会说找不到remote仓库的 项目名称.git 文件
	* Team --> Remote --> Push ,其中URI就是github上面指定的地址
	* source ref 选择 refs/heads/master  destination ref会自动填充，点击  Add Spec勾选Focus update开始提交
	* 如果有问题，可以选择提交的时候merge branch 'master' to github的项目地址

* 7.注意目录src下需要有一个main目录，否则运行时会报错：resource doesn't have a corresponding Go package.