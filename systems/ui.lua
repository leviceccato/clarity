local system = require('library.system')

local graphics = love.graphics

return function()
    local s = system({'appearance', 'position', 'ui'})

    s.draw = function()
        graphics.setColor(0, 0, 0, 0.4)
        graphics.rectangle('fill', 0, 0, 150, 35)
        graphics.setColor(1, 1, 1, 1)
        graphics.print('FPS: ' .. love.timer.getFPS(), 10, 10)
    end

    return s
end