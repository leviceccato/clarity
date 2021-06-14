local system = require('library.system')

local graphics = love.graphics
local timer = love.timer

return function(state)
    local s = system({})

    s.draw = function()
        graphics.setColor(0, 0, 0, 0.4)
        graphics.rectangle('fill', 0, 0, 150, 35)
        graphics.setColor(1, 1, 1, 1)
        graphics.print('FPS: ' .. timer.getFPS(), 10, 10)
    end

    return s
end