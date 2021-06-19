local graphics = love.graphics
local mouse = love.mouse
local min = math.min
local renderWidth = love.c.renderWidth
local renderHeight = love.c.renderHeight
local lastMouseX = 0
local lastMouseY = 0

local getWindowData = function()
    local windowWidth, windowHeight = graphics.getDimensions()
    return
        windowWidth,
        windowHeight
        min(windowWidth / renderWidth, windowHeight / renderHeight)
end

local getRawMousePosition = function()
    local mouseX, mouseY = mouse.getPosition()
    local windowWidth, windowHeight, scale = getWindowData()
    return
        (mouseX - (windowWidth - renderWidth * scale) * 0.5) / scale,
        (mouseY - (windowHeight - renderHeight * scale) * 0.5) / scale
end

local viewport = {}

viewport.getScale = function()
    local _, _, scale = getWindowData()
    return scale
end

viewport.getMousePosition = function()
    local mouseX, mouseY = getRawMousePosition()
    if mouseX >= 0 and mouseX <= renderWidth and mouseY >= 0 and mouseY <= renderHeight then
        lastMouseX = mouseX
        lastMouseY = mouseY
     end
     return lastMouseX, lastMouseY
end

viewport.beginRendering = function()
    local windowWidth, windowHeight, scale = getWindowData()

    graphics.push()
    graphics.translate((windowWidth - renderWidth * scale) * 0.5, (windowHeight - renderHeight * scale) * 0.5)
    graphics.scale(scale)

    return scale
end

viewport.endRendering = function()
    graphics.pop()

    local windowWidth, windowHeight, scale = getWindowData()
    local width = renderWidth * scale
    local height = renderHeight * scale

    graphics.setColor(0, 0, 0)
    graphics.rectangle("fill", 0, 0, windowWidth, 0.5 * (windowHeight - height))
    graphics.rectangle("fill", 0, windowHeight, windowWidth, -0.5 * (windowHeight - height))
    graphics.rectangle("fill", 0, 0,  0.5 * (windowWidth - width), windowHeight)
    graphics.rectangle("fill", windowWidth, 0, -0.5 * (windowWidth - width), windowHeight)
 end

return viewport