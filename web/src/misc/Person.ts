export enum Gender {
	MALE = 'male',
	FEMALE = 'female',
}

export interface Portrait {
	fileName: string;
	mimeType: string;
	dateTaken?: string;
}

export interface IPerson {
	id: number
	parentId: number
	name: string;
	path: string
	sex: string
	dateBirth: string;
	dateMarry: string;
	placeBirth: string;
	dateDeath: string;
	placeDeath: string;
	content: string;
	honor: string;
	children: Array<IPerson>;
}

export function newPerson(): IPerson {
	return {
		id: 0,
		parentId: 0,
		name: '',
		path: '',
		sex: '',
		dateBirth: '',
		dateMarry: '',
		placeBirth: '',
		dateDeath: '',
		placeDeath: '',
		content: '',
		honor: '',
		children: [],
	};
}