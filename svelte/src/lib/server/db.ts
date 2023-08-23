import { env } from "$env/dynamic/private"
import { MongoClient } from "mongodb"

export interface Paths {
	macos: string
	windows: string
	linux: string
}

export interface Dotfile {
	filename?: string
	username: string
	content: string
	paths: Paths
	isDirectory?: boolean
}

export const client = await MongoClient.connect(env.DB_CONN_STRING)
export const db = client.db("main")
export const dfColl = db.collection<Dotfile>("files")

// export interface Storer<T extends Document> {
// 	coll: Collection<T>
// }
