import { z } from "zod"
import { columnType } from "@/models/columnType.ts";
import { objectIdSchema } from "@/models/objectIdSchema.ts";

const boardType = z.object({
    ID: objectIdSchema,
    Name: z.string(),
    CreatedAt: z.date(),
    Columns: z.array(columnType)
})

export type boardModel = z.infer<typeof boardType>

