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

        local buttonImage = love.graphics.newImage('assets/sprites/title-button.png')
        local buttonSheet = json('assets/sprites/title-button.json')
        -- Align to middle
        local buttonX = (love.c.renderWidth / 2) - 50
        w.addEntity(buttonEntity({
            x = buttonX,
            y = 120,
            content = 'Start',
            image = buttonImage,
            sheet = buttonSheet,
            align = 'center',
            padding = 9,
            handler = function()
                state.activateWorld('main')
            end
        }))
        w.addEntity(buttonEntity({
            x = buttonX,
            y = 155,
            content = 'Quit',
            image = buttonImage,
            sheet = buttonSheet,
            align = 'center',
            padding = 9,
            handler = function()
                love.event.quit()
            end
        }))
 
        w.updateSystems()
        w.sortEntities()

        for index = 1, #w.systems do
            w.systems[index].load(arg)
        end
    end

    return w
end