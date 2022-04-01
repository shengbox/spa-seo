# 单页面应用SEO工具

解决单页面应用SEO问题

## 方案
![RUNOOB 图标](https://shengbox-picture-bed.oss-cn-hangzhou.aliyuncs.com/WechatIMG72.jpeg)  


1. 使用nginx常规部署spa服务给客户使. [nginx反向代理配置参考](#nginx反向代理配置参考)
2. 额外部署seo服务器，seo服务器内部使用chrome访问spa服务，渲染执行完毕的document给爬虫
```
注意点：在nginx配置上添加更加通过 user_agent 头来 判断是否是爬虫的配置，是爬虫就转向seo服务器
```

## docker 部署 seo服务
```dockerfile
docker run -p 8000:80 shengbox/spa-seo spa-seo -t https://www.spa.com -p 80
```

|  参数   | 必填  |说明 |
|  ----------  | ----  |---------- |
| 8000|true | 部署seo服务端口,可自行修改  |
|  80 | true  |  docker内部go服务的端口,可自行修改，但前后必须保持一致  |
|  -t  | false |  spa客户端域名地址， 默认获取nginx代理过来的域名，<br/> 如果填写了参数例如https://www.spa.com, 内部设置的该参数为客户端域名地址   |






## 部署spa服务 

### nginx反向代理配置参考

```
# seo：服务端动态渲染方案-nginx判断爬虫

server {
        listen 80;
        add_header Access-Control-Allow-Origin *;
        # gzip config
        client_max_body_size 100M;
        gzip on;
        gzip_min_length 1k;
        gzip_comp_level 9;
        gzip_types text/plain text/css text/javascript application/json application/javascript application/x-javascript application/xml;
        gzip_vary on;
        gzip_disable "MSIE [1-6]\.";
        root /usr/share/nginx/html;

        location ~.*\.html$ {
            add_header Cache-Control "no-cache, no-store";
        }

        location / {
            # 用于是否是判断爬虫
            if ($http_user_agent ~* "googlebot|google-structured-data-testing-tool|Mediapartners-Google|bingbot|linkedinbot|baiduspider|360Spider|Sogou Spider|Yahoo! Slurp China|Yahoo! Slurp|twitterbot|facebookexternalhit|rogerbot|embedly|quora link preview|showyoubot|outbrain|pinterest|slackbot|vkShare|W3C_Validator") {
                set $agent $http_user_agent;
                 # 代理到seo服务器
                proxy_pass  http://www.seo.com;
                break;
            }
            # 不是就正常访问spa服务器
            # 用于配合 browserHistory使用
            try_files $uri $uri/ /index.html;

            # 如果有资源，建议使用 https + http2，配合按需加载可以获得更好的体验
            # rewrite ^/(.*)$ https://preview.pro.ant.design/$1 permanent;
        }
        # spa服务请求后端api代理（可选）
        location  /api/ {
            proxy_pass  https://api.xx.com/api/;
        }
    }
}
```

