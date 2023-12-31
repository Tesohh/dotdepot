import type { Handler, HandlerEvent, HandlerContext } from "@netlify/functions"
import { MongoClient } from "mongodb"
import bcrypt from "bcrypt"

import dotenv from "dotenv"
import sortKeysRecursive from "sort-keys-recursive"
dotenv.config()

const client = MongoClient.connect(process.env.DB_CONN_STRING ?? "")

const handler: Handler = async (event: HandlerEvent, context: HandlerContext) => {
	const db = (await client).db("main")
	const dfColl = db.collection(event.queryStringParameters?.collection ?? "")
	const userColl = db.collection("users")

	const username = event.queryStringParameters?.username
	if (username == undefined) return { statusCode: 400, body: "username missing" }
	const password = event.queryStringParameters?.password
	if (password == undefined) return { statusCode: 400, body: "password missing" }
	if (event.body == null) return { statusCode: 400, body: "missing body" }

	const user = await userColl.findOne({ username: username })
	if (user == null) return { statusCode: 404, body: "user not found" }
	if (!bcrypt.compareSync(password, user.password ?? ""))
		return { statusCode: 401, body: "wrong password" }

	const jquery = JSON.parse(event.body || "") as object
	const sorted = sortKeysRecursive(jquery)

	const res = await dfColl.insertOne(sorted)
	console.log(res.acknowledged)
	console.log(event.body)

	return {
		statusCode: 200,
		insertedId: res.insertedId
	}
}

export { handler }
