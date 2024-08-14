import { z } from "zod";
import { columnSchema } from "@/models/columnType.ts";
import { ObjectIdSchema } from "@/models/objectIdSchema.ts";

export const boardSchema = z.object({
    ID: ObjectIdSchema,
    Name: z.string(),
    CreatedAt: z.date(),
    Columns: z.array(columnSchema)
});

export type boardType = z.infer<typeof boardSchema>;