import type { Handler, HandlerEvent, HandlerContext } from "@netlify/functions"
import { MongoClient } from "mongodb"
import bcrypt from "bcrypt"

import dotenv from "dotenv"
dotenv.config()

const client = MongoClient.connect(process.env.DB_CONN_STRING ?? "")

const handler: Handler = async (event: HandlerEvent, context: HandlerContext) => {
	const db = (await client).db("main")
	const userColl = db.collection("users")

	const username = event.queryStringParameters?.username
	if (username == undefined) return { statusCode: 400, body: "username missing" }
	const password = event.queryStringParameters?.password
	if (password == undefined) return { statusCode: 400, body: "password missing" }

	const user = await userColl.findOne({ username: username })
	if (user != null) return { statusCode: 404, body: "user already exists" }

	const hash = bcrypt.hashSync(password, 8)
	const res = await userColl.insertOne({ username: username, password: hash })

	return {
		statusCode: 200,
		insertedId: res.insertedId
	}
}

export { handler }
