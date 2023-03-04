import asyncio
import logging
import sys
from typing import Optional

import httpx

TODO_SERVER = "https://jsonplaceholder.typicode.com/todos"

logger = logging.getLogger(__name__)


async def get_todos(client: httpx.AsyncClient, index: int) -> Optional[str]:
    try:
        await asyncio.sleep(0.5)
        response = await client.get(f"{TODO_SERVER}/{index}")
        response.raise_for_status()
    except httpx.HTTPStatusError as exc:
        logger.warning(f"Failed to get todo {index}: {exc}")
        return
    return response.json()["title"]


async def main(num_todos: int) -> None:
    async with httpx.AsyncClient() as client:
        tasks = [get_todos(client, index + 1) for index in range(num_todos)]
        results = await asyncio.gather(*tasks)

    print("# Todos")
    for todo in results:
        if todo is None:
            continue
        print(f"* {todo}")


if __name__ == "__main__":
    asyncio.run(main(int(sys.argv[1])))
