# Golang 日志处理

## github的开源日志库

- logrus：https://github.com/sirupsen/logrus
- zap：https://github.com/uber-go/zap
  
Zap提供了可插拔的Hooks机制，允许开发者在生成日志信息前或后执行自定义的代码逻辑，从而实现了更加灵活的扩展

- zerolog：https://github.com/rs/zerolog

与传统的文本日志不同，zerolog允许用户以结构化的方式记录日志，这样可以更方便地进行搜索和分析

- seelog：https://github.com/cihub/seelog

seelog支持使用XML或JSON格式配置日志记录器，支持多种日志记录方式，如文件、网络、邮件等；支持根据日志级别、日志来源、时间戳等多个维度进行筛选和过滤；seelog提供了插件机制，允许用户开发自己的插件以扩展其功能

- go-kit/log：https://github.com/go-kit/log

支持将日志按照不同的层次进行分类输出，比如将Web层、业务逻辑层和数据访问层的日志分别输出到不同的文件中。