local json = require('utilities.json')
local world = require('library.world')

local drawSystem = require('systems.draw')
local animationSystem = require('systems.animation')
local debugSystem = require('systems.debug')
local hoverSystem = require('systems.hover')
local clickSystem = require('systems.click')
local eventSystem = require('systems.event')
local inputSystem = require('systems.input')

local titleBgEntity = require('entities.title-bg')
local buttonEntity = require('entities.button')

return function(state)
    local w = world('title')

    w.load = function(arg)
        w.addSystem(drawSystem(state))
        w.addSystem(animationSystem(state))
        w.addSystem(debugSystem(state))
        w.addSystem(hoverSystem(state))
        w.addSystem(clickSystem(state))
        w.addSystem(eventSystem(state))
        w.addSystem(inputSystem(state))

        w.addEntity(titleBgEntity())

        local buttonImage = love.graphics.newImage('assets/sprites/title-button.png')
        local buttonSheet = json('assets/sprites/title-button.json')
        local currentFont = love.graphics.getFont()
        -- Align to middle
        local buttonX = (love.c.renderWidth / 2) - 50

        local startButtonString = 'START'
        local startButtonText = love.graphics.newText(currentFont, startButtonString)
        w.addEntity(buttonEntity({
            x = buttonX,
            y = 120,
            content = startButtonString,
            image = buttonImage,
            sheet = buttonSheet,
            align = 'center',
            padding = 9,
            event = {'activate-world', {name = 'main'}}
        }))

        local quitButtonString = 'QUIT'
        local quitButtonText = love.graphics.newText(currentFont, quitButtonString)
        w.addEntity(buttonEntity({
            x = buttonX,
            y = 155,
            content = 'QUIT',
            image = buttonImage,
            sheet = buttonSheet,
            align = 'center',
            padding = 9,
            event = {'quit'}
        }))

        w.updateSystems()
        w.sortEntities()

        for index = 1, #w.systems do
            w.systems[index].load(arg)
        end
    end

    return w
end