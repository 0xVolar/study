# TS的值和类型

## 三种基本类型

- **any类型**：any类型代表了所有类型，any类型可以兼容所有类型，同时也可以赋值给所有类型，这就导致可能会产生**污染问题**。建议很少使用此类型
  - 使用编译命令时可以使用`tsc --noImplicitAny xx.ts`避免any的使用。
- **unkonwn类型**：是一个较为严格一点的any类型，unknown类型的值只能够赋给any型和unknown型。**不能直接调用unknown型的方法和属性**。unknown类型只能够进行**比较运算**、**取反运算**、**typeof**、**instanceof**。
- **never类型**：never类型不代表任何一种类型，可以抽象的理解为一个空集，同时never的类型的值可以赋给任何一个类型。

any类型和unknown类型都是顶层类型，而never类型是底层类型。
