package main

func drawMap() string {
	output := ""
	for y := 0; y < mapHeight; y++ {
		for x := 0; x < mapWidth; x++ {
			if x == playerX && y == playerY {
				output += string(playerChar)
			} else {
				output += string(gameMap[y][x])
			}
		}
		output += "\n"
	}
	return output
}
