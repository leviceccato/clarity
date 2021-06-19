local graphics = love.graphics
local mouse = love.mouse
local renderWidth = love.c.renderWidth
local renderHeight = love.c.renderHeight
local min = math.min
local floor = math.floor

local lastMouseX = 0
local lastMouseY = 0

local viewport = {}

viewport.getWindowData = function()
    local windowWidth, windowHeight = graphics.getDimensions()
    return
        windowWidth,
        windowHeight,
        floor(min(windowWidth / renderWidth, windowHeight / renderHeight))
end

local getRawMousePosition = function()
    local mouseX, mouseY = mouse.getPosition()
    local windowWidth, windowHeight, scale = viewport.getWindowData()
    return
        (mouseX - (windowWidth - renderWidth * scale) * 0.5) / scale,
        (mouseY - (windowHeight - renderHeight * scale) * 0.5) / scale
end

viewport.getMousePosition = function()
    local mouseX, mouseY = getRawMousePosition()
    if mouseX >= 0 and mouseX <= renderWidth and mouseY >= 0 and mouseY <= renderHeight then
        lastMouseX = mouseX
        lastMouseY = mouseY
     end
     return lastMouseX, lastMouseY
end

return viewport