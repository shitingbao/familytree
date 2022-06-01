# look and study

## postgresql

## golang 1.18  

1.gin；2.gorm

## vue3  

1.vite ；2.ants;3.less

## study address

1.[vue3](https://vue3.chengpeiquan.com/communication.html#%E5%88%9B%E5%BB%BA%E5%92%8C%E7%A7%BB%E9%99%A4%E7%9B%91%E5%90%AC%E4%BA%8B%E4%BB%B6-new)

2.[vite](https://vitejs.cn/guide/)

3.[flex](https://www.bbsmax.com/A/l1dyQD7nde/)

4.[ants](https://www.antdv.com/components/button-cn)

5.[gin](https://github.com/gin-gonic/gin)

## start

1. go后端到 familytree/family/cmd 目录下，执行 go mod tidy (下载包)，go build 构建执行文件，执行生成的 ./cmd  

2. web前端 familytree/familyweb 目录下，执行 yarn 安装依赖包，yarn dev 启动项目  

## Project design idea

### 数据库表设计

1.通过一张表存储树结构;  
2.单行存储父节点，以及完整的根目录到本节点的路径;  
3.父节点可快速构建树;
4.完整路径可一次性获取出所有树节点，并能获取父节点 id，相当于双向链表;  
5.配偶节点的 marryid 存储了对应树配偶节点的 id  ，其根目录和配偶 id 相同，但没有父节点id

### vue 节点显示

1.构建树结构模型  
2.组件根据结构模型递归显示树结构
