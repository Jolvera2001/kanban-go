import { z } from "zod";
import { ObjectId } from "bson";

export const ObjectIdSchema = z.custom<ObjectId>((value) => {
    return ObjectId.isValid(value)
}, {
    message: "Invalid Object Id"
});