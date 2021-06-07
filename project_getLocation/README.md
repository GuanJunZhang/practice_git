### 代码解构

- controller web接口,与http相关的代码都要放在这里,如参数合法性、格式转换等
- service 业务逻辑层,处于应用环境(如web、grpc等)与model之间,负责复杂业务逻辑处理. 如业务逻辑十分简单,可省略service层,直接在model层实现. *它不应关心运行环境*
- model 负责数据库entity定义,及简单数据库操作.若涉及多表操作或较复杂的业务逻辑,应放在service层
- model/model_event.go 可集中定义model中事件相关业务逻辑,如增加评论后增加反馈的评论计数这种非严格强制性业务逻辑(失败也不影响功能)
- conf 配置生成逻辑
- route web路由配置
- utils 工具类,负责非数据库entity定义
- static 静态资源目录，包括Js，css，jpg等等，可以通过echo框架配置，直接让用户访问
- views 视图模板目录，存放各个模块的视图模板，当然有些项目只有api，是不需要视图部分，可以忽略这个目录
