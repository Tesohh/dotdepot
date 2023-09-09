import type { Handler, HandlerEvent, HandlerContext } from "@netlify/functions"
import { MongoClient } from "mongodb"

import dotenv from "dotenv"
dotenv.config()

const client = MongoClient.connect(process.env.DB_CONN_STRING ?? "")

const handler: Handler = async (event: HandlerEvent, context: HandlerContext) => {
	const db = (await client).db("main")
	const dfColl = db.collection(event.queryStringParameters?.collection ?? "")

	return {
		statusCode: 200,
		body: JSON.stringify(await dfColl.find(JSON.parse(event.body || "")).toArray())
	}
}

export { handler }
