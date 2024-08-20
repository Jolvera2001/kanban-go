import {boardType} from "@/models/boardType.ts";
import {ObjectId} from "bson";

export interface IBoardService {
    createBoard(boardName: string): Promise<boardType>;
    getBoards(): Promise<boardType[]>;
    updateBoard(updatedBoard: boardType): Promise<string>;
    deleteBoard(id: ObjectId): Promise<string>;
}