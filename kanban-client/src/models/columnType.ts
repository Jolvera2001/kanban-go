import { z } from "zod"
import { objectIdSchema } from "@/models/objectIdSchema.ts";


export const columnType = z.object({
    ID: objectIdSchema,
    Name: z.string(),
    CreatedAt: z.date()
})

export type columnModel = z.infer<typeof columnType>