import type { Handler, HandlerContext, HandlerEvent } from "@netlify/functions"
import { MongoClient } from "mongodb"

import dotenv from "dotenv"
import sortKeysRecursive from "sort-keys-recursive"
dotenv.config()

const client = MongoClient.connect(process.env.DB_CONN_STRING ?? "")

const handler: Handler = async (event: HandlerEvent, context: HandlerContext) => {
	const db = (await client).db("main")
	const dfColl = db.collection(event.queryStringParameters?.collection ?? "")
	const jquery = JSON.parse(event.body || "") as object
	const sorted = sortKeysRecursive(jquery)

	const df = await dfColl.findOne(sorted)
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
