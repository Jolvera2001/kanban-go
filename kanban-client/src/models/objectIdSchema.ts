import { z } from "zod";
import { ObjectId } from "bson";

export const objectIdSchema = z.custom<ObjectId>((value) => {
    return ObjectId.isValid(value)
}, {
    message: "Invalid MongoDB ObjectId"
})