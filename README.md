# 单页面应用SEO工具

解决单页面应用SEO问题

```
# seo2：服务端动态渲染方案-nginx判断爬虫
server {
    listen  80;
    server_name www.cmvalue.seo2;

    add_header Access-Control-Allow-Origin *;

    location /homepage {
        if ($http_user_agent ~* "googlebot|google-structured-data-testing-tool|Mediapartners-Google|bingbot|linkedinbot|baiduspider|360Spider|Sogou Spider|Yahoo! Slurp China|Yahoo! Slurp|twitterbot|facebookexternalhit|rogerbot|embedly|quora link preview|showyoubot|outbrain|pinterest|slackbot|vkShare|W3C_Validator") {
          set $agent $http_user_agent;
          proxy_pass  http://www.cmvalue.seo2:7001;
          break;
        }

        alias /var/www/f2e/yl-homepage/;
        index  index.html index.htm;
        try_files $uri $uri/ /homepage/index.html;
    }

    location ~ /api/ {
        proxy_pass  https://www.cmvalue.com;
    }
}
```

## docker 部署

docker run -p 8000:80 shengbox/spa-seo spa-seo -t https://www.cmvalue.com -p 80

-t 为单页面应用域名
-p 为本服务端口