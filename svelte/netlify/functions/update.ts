import type { Handler, HandlerEvent, HandlerContext } from "@netlify/functions"
import { MongoClient } from "mongodb"
import bcrypt from "bcrypt"

import dotenv from "dotenv"
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

	const body = JSON.parse(event.body)
	const formaterr = "incorrect body! format: {query: <the query>, doc: <the new doc>}"
	if (body.query == undefined) return { statusCode: 400, body: formaterr }
	if (body.doc == undefined) return { statusCode: 400, body: formaterr }

	const user = await userColl.findOne({ username: username })
	if (user == null) return { statusCode: 404, body: "user not found" }
	if (!bcrypt.compareSync(password, user.password ?? ""))
		return { statusCode: 401, body: "wrong password" }

	const res = await dfColl.findOneAndUpdate(body.query, { $set: body.doc })

	return {
		statusCode: res.ok ? 200 : 400
	}
}

export { handler }
