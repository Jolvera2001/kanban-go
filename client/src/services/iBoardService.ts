import {boardType} from "@/models/boardType.ts";
import {ObjectId} from "bson";

export interface IBoardService {
    createBoard(boardName: string): Promise<boardType>;
    getBoards(): Promise<boardType[]>;
    getBoardById(id: ObjectId): Promise<boardType>;
    updateBoard(updatedBoard: boardType): string;
    deleteBoard(id: ObjectId): string;
}