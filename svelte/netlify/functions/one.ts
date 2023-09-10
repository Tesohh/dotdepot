import type { Handler, HandlerEvent, HandlerContext } from "@netlify/functions"
import { MongoClient } from "mongodb"

import dotenv from "dotenv"
dotenv.config()

const client = MongoClient.connect(process.env.DB_CONN_STRING ?? "")

const handler: Handler = async (event: HandlerEvent, context: HandlerContext) => {
	const db = (await client).db("main")
	const dfColl = db.collection(event.queryStringParameters?.collection ?? "")
	const df = await dfColl.findOne(JSON.parse(event.body || ""))
	if (df == null) {
		return {
			statusCode: 404,
			body: "couldn't find doc"
		}
	}
	return {
		statusCode: 200,
		body: JSON.stringify(df)
	}
}

export { handler }
