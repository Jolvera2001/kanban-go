import { boardType } from "@/models/boardType";
import { IBoardService } from "@/services/iBoardService.ts";
import { ObjectId } from "bson";

export class BoardService implements IBoardService {

    createBoard(boardName: string): Promise<boardType[]> {
        throw new Error("Method not implemented.");
    }
    getBoards(): Promise<boardType[]> {
        throw new Error("Method not implemented.");
    }
    getBoardById(id: ObjectId): Promise<boardType> {
        throw new Error("Method not implemented.");
    }
    updateBoard(updatedBoard: boardType): string {
        throw new Error("Method not implemented.");
    }
    deleteBoard(id: ObjectId): string {
        throw new Error("Method not implemented.");
    }

}