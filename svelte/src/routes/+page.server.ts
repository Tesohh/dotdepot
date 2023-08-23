import { dfColl, type Dotfile } from "$lib/server/db"
import { error } from "@sveltejs/kit"
import type { PageServerLoad } from "./$types"

export const load: PageServerLoad = async () => {
	const docs = await dfColl.find({ username: "tesohhhh" }, { projection: { _id: 0 } }).toArray()
	if (docs == null) {
		throw error(404)
	}
	return { docs }
}
