local system = require('library.system')

local graphics = love.graphics
local timer = love.timer

return function(state)
    local s = system({'appearance', 'position', 'ui'})

    s.draw = function()
        for index = 1, #s.entities do
            local e = s.entities[index]
            graphics.draw(
                e.appearance.image,
                e.animation and e.animation.getFrame() or e.appearance.frame,
                e.position.x,
                e.position.y,
                0, -- rotation
                2, 2 -- scale
            )
        end
        graphics.setColor(0, 0, 0, 0.4)
        graphics.rectangle('fill', 0, 0, 150, 35)
        graphics.setColor(1, 1, 1, 1)
        graphics.print('FPS: ' .. timer.getFPS(), 10, 10)
    end

    return s
end