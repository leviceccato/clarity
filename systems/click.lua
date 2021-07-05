local system = require('library.system')
local viewport = require('utilities.viewport')

local eventEntity = require('entities.event')

local mouse = love.mouse

return function(state)
    local s = system({'click', 'position', 'appearance'})

    local isWithin = function(point, e)
        local width = e.appearance.width
        local height = e.appearance.height
        if e.animation then
            local _, _, w, h = e.animation.getFrame():getViewport()
            width = w
            height = h
        end
        local isWithinX = point.x >= e.position.x and point.x <= e.position.x + width
        local isWithinY = point.y >= e.position.y and point.y <= e.position.y + height
        return isWithinX and isWithinY
    end

    s.update = function(dt)
        local e
        for index = 1, #s.entities do
            e = s.entities[index]
            
            local clickData = e.click.data
            local clickInput = state.inputs.click

            -- No clicks are happening or have happened
            if not clickInput and not clickData then
                goto done
            end
            
            -- Mouse was released after click
            if not clickInput and clickData then
                -- If it's in the right place run the handler
                local x, y = viewport.getMousePosition()
                if isWithin({x = x, y = y}, e) then
                    state.activeWorld.addEntity(eventEntity(e.click.event), true)
                end
                e.click.data = nil
                goto done
            end
            
            -- Otherwise mouse was clicked
            if clickInput then
                -- If it's in the right place assign the data to the component
                if isWithin(clickInput, e) then
                    e.click.data = clickInput
                end
            end

            ::done::
        end
    end

    return s
end