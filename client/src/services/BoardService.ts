import {boardSchema, boardType} from "@/models/boardType";
import { IBoardService } from "@/services/iBoardService.ts";
import { ObjectId } from "bson";

export class BoardService implements IBoardService {
    private readonly baseApiUrl : string;

    constructor() {
        this.baseApiUrl = "http://localhost:8080/api/v1/"
    }

    async createBoard(boardName: string): Promise<boardType> {
        const data = { Name: boardName}

        try {
            const response = await fetch(`${this.baseApiUrl}/board`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data),
            });

            if (!response.ok) {
                throw new Error("Network response was not ok")
            }

            const json = await response.json();
            return boardSchema.parse(json);

        } catch (e) {
            console.log(e);
            return {} as boardType;
        }
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

export class BoardServiceFactory {
    public static createService(): IBoardService {
        return new BoardService();
    }
}