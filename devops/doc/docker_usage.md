
``` bash

根据dockerfile创建镜像
docker build -t ali-proxy-xgo .

启动容器，进入容器bash 查看变量等内容  xgo不适合进入。
docker run -it --rm --name=ali-proxy-xgo ali-proxy-xgo /bin/bash 

推送到docker hub 
docker tag ali-proxy-xgo youwen21/ali-proxy-xgo
docker push youwen21/ali-proxy-xgo
```