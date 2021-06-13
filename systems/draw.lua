local system = require('library.system')

local graphics = love.graphics

return function()
    local s = system({'appearance', 'position'})

    s.draw = function()
        for index = 1, #s.entities do
            local e = s.entities[index]
            if e.animation then
                graphics.draw(
                    e.appearance.image,
                    e.animation.getQuad,
                    e.position.x,
                    e.position.y
                )
            else
                graphics.draw(
                    e.appearance.image,
                    e.position.x,
                    e.position.y
                )
            end
        end
    end

    return s
end