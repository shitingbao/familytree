// 基本单链表的树结构结构
export class Member {
	id = 0;
	parentId = 0;
	name = '';
	path = '';
	sex = '';
	dateBirth = '';
	dateMarry = '';
	placeBirth = '';
	dateDeath = '';
	placeDeath = '';
	content = '';
	honor = '';
	familySimple = ''
	marryId = 0
	memberChildren: Member[] = [];
	marryMember: Member = {} as Member;

	initData(row: any) {
		// console.log("rwo:", row)
		this.id = row?.id;
		this.parentId = row?.parentId;
		this.name = row?.name;
		this.path = row?.path;
		this.sex = row?.sex == 1 ? '男' : '女';
		this.dateBirth = row?.dateBirth;
		this.dateMarry = row?.dateMarry;
		this.placeBirth = row?.placeBirth;
		this.dateDeath = row?.dateDeath;
		this.placeDeath = row?.placeDeath;
		this.content = row?.content;
		this.honor = row?.honor;
		this.familySimple = row?.familySimple;
		// this.children = row?.children
	}

	// 构造头节点
	getHeader(rows: any[]) {
		if (rows.length == 0) {
			return
		}
		var root = rows.find((e) => {
			return e.parentId == 0
		})
		this.initData(root)
		this.construct(this, rows)
	}

	// 从反馈的所有节点里，构建基本树结构
	// 递归找到自己所有孩子
	// 同时加入配偶节点
	construct(root: Member, rows: any[]) {
		var rs = rows.filter((item) => {
			return item.parentId == root.id
		})
		if (rs.length > 0) {
			rs.forEach((e) => {
				// console.log("for root:", root)
				var node = new Member()
				node.initData(e)
				root.memberChildren.push(node)
				this.construct(node, rows)
			})
		}
		// 加入配偶对象
		var marryNode = rows.find((item) => {
			return item.marryId == root.id
		})

		var n = new Member()
		n.initData(marryNode)
		root.marryMember = n

		console.log("marryNode:", marryNode, this.marryMember)
	}
}
