# 文件操作封装


- 文件上传下载
  - 腾讯云
  - 阿里云
  - 华为云


## 配置

- 每个云单独配置
- mfile仅选择一个云进行访问，通过配置确定


## 提供功能

- 上传文件
- 加密上传文件
- 加密下载文件



## 使用说明

```

配置


FileServiceMode   文件服务模式  tencent  aliyun huaweicloud  支持三种

每种服务模式有其自己的配置


腾讯云

TencentCloudCosSecretID    应用key
TencentCloudCosSecretKey   加密key
TencentCloudCosBucketUrl   桶url
TencentCloudCosServiceUrl  加速url
TencentCloudCosCustomerKey 加密key 32位


阿里云

AliyunOssAccessKeyId     应用key
AliyunOssAccessKeySecret 加密key
AliyunOssEndpoint        上传url
AliyunOssServiceUrl      加速url
AliyunOssBucket          桶


华为云




```










