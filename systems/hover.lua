local system = require('library.system')
local viewport = require('utilities.viewport')

local mouse = love.mouse

return function(state)
    local s = system({'hover', 'position', 'appearance'})

    s.update = function(dt)
        local mouseX, mouseY = viewport.getMousePosition()
        local e
        for index = 1, #s.entities do
            e = s.entities[index]
            local width = e.appearance.width
            local height = e.appearance.height
            if e.animation then
                local _, _, w, h = e.animation.getFrame():getViewport()
                width = w
                height = h
            end
            local isWithinX = mouseX >= e.position.x and mouseX <= e.position.x + width
            local isWithinY = mouseY >= e.position.y and mouseY <= e.position.y + height
            local isHovered = isWithinX and isWithinY
            local hasHoverChanged = isHovered ~= e.hover.isHovered
            if hasHoverChanged then
                e.hover.isHovered = isHovered
                if e.animation and e.animation.sequences['hover'] then
                    if isHovered then
                        e.animation.previousSequence = e.animation.sequence
                        e.animation.sequence = 'hover'
                    else
                        e.animation.sequence = e.animation.previousSequence
                        e.animation.previousSequence = 'hover'
                    end
                end
            end
        end
    end

    return s
end