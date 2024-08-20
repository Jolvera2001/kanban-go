import { ObjectIdSchema } from "@/models/objectIdSchema.ts";
import { z } from "zod";

export const columnSchema = z.object({
    ID: ObjectIdSchema,
    Name: z.string(),
    CreatedAt: z.date()
});

export type columnType = z.infer<typeof columnSchema>;