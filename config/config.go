package config

//加载配置信息的代码需要优先于其他代码执行
// 所以要在main 包的init里面去执行下面的函数
// 这个是个空函数,不做任何处理,调用他仅仅是为了引入包,从而触发包内各个init方法的执行
func Initialize() {

}
