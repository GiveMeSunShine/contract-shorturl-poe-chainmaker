本项目是一个存证合约，主要业务未存储短链接相关信息

------------

### 准备合约编译环境：

```bash
cd $GOPATH/src/
docker pull chainmakerofficial/chainmaker-go-contract:2.0.0
docker run -it --name chainmaker-go-contract -v `pwd`/contract-shorturl-chainmaker:/home/contract-shorturl-chainmaker chainmakerofficial/chainmaker-go-contract:2.0.0 bash
# 或者先后台启动
docker run -d  --name chainmaker-go-contract -v `pwd`/contract-shorturl-poe-chainmaker:/home/contract-shorturl-poe-chainmaker chainmakerofficial/chainmaker-go-contract:2.0.0 bash -c "while true; do echo hello world; sleep 5;done"
# 再进入容器
docker exec -it chainmaker-go-contract bash
```
### 编译合约
```bash
root@3b2a6218126d:/# ls 
bin  boot  data  dev  etc  home  lib  lib64  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
root@3b2a6218126d:/# cd /home/contract-shorturl-poe-chainmaker/
root@3b2a6218126d:/home/contract-shorturl-poe-chainmaker# ls
Makefile  build.sh  chainmaker.go  chainmaker_rs.go  convert  easycodec.go  go.mod  main.go
root@3b2a6218126d:/home/contract-shorturl-poe-chainmaker# ./build.sh
root@3b2a6218126d:/home/contract-shorturl-poe-chainmaker# ls
Makefile  build.sh  chainmaker.go  chainmaker_rs.go  convert  easycodec.go  go.mod  main.go  shorturl-poe-chainmaker-contract-go-2.0.0.wasm

```

### 合约部署
```bash
./cmc client contract user create \
--contract-name=shorturl \
--runtime-type=GASM \
--byte-code-path=/home/ysh/workspace/goPro/src/chainmaker-sdk-go-demo/contract/wasm/shorturl-poe-chainmaker-contract-go-2.0.0.wasm \
--version=1.0 \
--sdk-conf-path=./testdata/sdk_config.yml \
--admin-key-file-paths=./testdata/crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.tls.key,./testdata/crypto-config/wx-org2.chainmaker.org/user/admin1/admin1.tls.key,./testdata/crypto-config/wx-org3.chainmaker.org/user/admin1/admin1.tls.key \
--admin-crt-file-paths=./testdata/crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.tls.crt,./testdata/crypto-config/wx-org2.chainmaker.org/user/admin1/admin1.tls.crt,./testdata/crypto-config/wx-org3.chainmaker.org/user/admin1/admin1.tls.crt \
--sync-result=true \
--params="{}"
```
```bash
response: message:"OK" contract_result:<result:"\n\010shorturl\022\0031.0\030\004*<\n\026wx-org1.chainmaker.org\020\001\032 \337Nf:\220\257\026/\266\277d\250\027\371\247E\204q\217u)\300\363\003\264\035\332\3348\244~\335" message:"OK" > tx_id:"2a47b53faa3a4361a4d1b56f9bca86573a8f4e5fd0cd472cb8d75c4f464dd33c" 
```
### 合约调用
#### 存储
```bash
存储命令
./cmc client contract user invoke \
--contract-name=shorturl \
--method=save \
--sdk-conf-path=./testdata/sdk_config.yml \
--params="{\"code\":\"mfpIiQcng\",\"long_url\":\"https://docs.chainmaker.org.cn/v2.0.0/html/dev/%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%B7%A5%E5%85%B7.html\",\"short_url\":\"http://uptc.umpay.com/stl/mfpIiQcng\",\"description\":\"测试数据\",\"creator\":\"admin\",\"version\":\"1.0.0\",\"time\":\"1637130881\"}" \
--sync-result=true
响应结果：
INVOKE contract resp, [code:0]/[msg:OK]/[contractResult:result:"mfpIiQcng" gas_used:12873433 contract_event:<topic:"topic_short" tx_id:"0bcd39d87e0a44b6bc46b7abc135c124d26251827c444d6bb58ea0d06c20131c" contract_name:"shorturl" contract_version:"1.0" event_data:"mfpIiQcng" > ]/[txId:0bcd39d87e0a44b6bc46b7abc135c124d26251827c444d6bb58ea0d06c20131c]
```
#### 查询
```bash
查询命令：
./cmc client contract user invoke \
--contract-name=shorturl \
--method=find_by_code \
--sdk-conf-path=./testdata/sdk_config.yml \
--params="{\"code\":\"mfpIiQcng\"}" \
--sync-result=true

响应结果：
INVOKE contract resp, [code:0]/[msg:OK]/[contractResult:result:"{\"shortUrl\":\"http://uptc.umpay.com/stl/mfpIiQcng\",\"longUrl\":\"https://docs.chainmaker.org.cn/v2.0.0/html/dev/%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%B7%A5%E5%85%B7.html\",\"code\":\"mfpIiQcng\",\"description\":\"\346\265\213\350\257\225\346\225\260\346\215\256\",\"creator\":\"admin\",\"version\":\"1.0.0\",\"time\":1637130881}" gas_used:8460460 ]/[txId:14248c745f1642789dd1031bac7745a924bb7f4064654a79ae709223eadafa0f]

```



