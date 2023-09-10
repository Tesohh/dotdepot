import { dfColl, type Dotfile } from "$lib/server/db"
import { error } from "@sveltejs/kit"
import type { PageServerLoad } from "./$types"

export const load: PageServerLoad = async ({ params }) => {
	let docs = await dfColl.find({ username: params.username }).toArray()
	if (docs == null) {
		throw error(404)
	}
	let dfs: Dotfile[] = []
	for (let i = 0; i < docs.length; i++) {
		dfs.push(docs[i])
		dfs[i].safeID = docs[i]._id.toString()

		// @ts-expect-error 2339
		delete dfs[i]._id
	}
	return { df: dfs.find((v) => v.safeID == params.dotfileid), depotname: params.username }
}
