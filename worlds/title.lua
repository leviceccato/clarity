local json = require('utilities.json')

local world = require('library.world')
local drawSystem = require('systems.draw')
local animationSystem = require('systems.animation')
local debugSystem = require('systems.debug')
local hoverSystem = require('systems.hover')
local titleBgEntity = require('entities.title-bg')
local buttonEntity = require('entities.button')

return function(state)
    local w = world('title')

    w.load = function(arg)
        w.addSystem(drawSystem(state))
        w.addSystem(animationSystem(state))
        w.addSystem(debugSystem(state))
        w.addSystem(hoverSystem(state))

        w.addEntity(titleBgEntity())

        -- Buttons
        local image = love.graphics.newImage('assets/sprites/title-button.png')
        local sheet = json('assets/sprites/title-button.json')
        -- Align to middle
        local buttonX = (love.c.renderWidth / 2) - 50
        w.addEntity(buttonEntity(
            buttonX, 120, 'Start', image, sheet,
            'center', 9, {215 / 255, 196 / 255, 91 / 255, 1}
        ))
        w.addEntity(buttonEntity(
            buttonX, 155, 'Quit', image, sheet,
            'center', 9, {215 / 255, 196 / 255, 91 / 255, 1}
        ))

        w.updateSystems()
        w.sortEntities()

        for index = 1, #w.systems do
            w.systems[index].load(arg)
        end
    end

    return w
end