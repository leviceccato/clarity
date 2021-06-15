local system = require('library.system')

local graphics = love.graphics

return function(state)
    local s = system({'appearance', 'position'})

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
    end

    return s
end