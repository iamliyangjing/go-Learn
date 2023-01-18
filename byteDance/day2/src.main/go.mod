module example/project/app  // 依赖管理基本单元

go 1.1.16  // 原生库

require (
	example/lib1 v1.0.2
	example/lib2 v1.0.1 //indirect
	example/lib3 v0.0.0
)
// 依赖标识 ：[Module Path][Version/Pseudo-version]