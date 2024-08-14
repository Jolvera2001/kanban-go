import { useState } from 'react'
import {boardType} from "@/models/boardType.ts";
import {Avatar, AvatarFallback, AvatarImage} from "@/components/ui/avatar.tsx";
import {Button} from "@/components/ui/button.tsx";
import {Separator} from "@/components/ui/separator.tsx";

function App() {
  const [boardList, setBoardList] = useState<boardType[] | null>(null);
  const [currBoard, setCurrBoard] = useState<boardType | null>(null);

  return (
    <>
        <div className="">
            <div className="flex flex-row h-5/6 my-2 space-x-2">
                <div className="flex flex-col p-2 h-full">
                    <Avatar>
                        <AvatarImage />
                        <AvatarFallback>TS</AvatarFallback>
                    </Avatar>
                </div>
                <Separator orientation="vertical"/>
                <div className="flex flex-col p-2 w-5/6 text-2xl">
                    <div className="flex flex-row justify-normal mb-3">
                        <h1>Board Name Here</h1>
                    </div>
                    <Separator className="w-full"/>
                </div>
            </div>
        </div>
    </>
  )
}

export default App
