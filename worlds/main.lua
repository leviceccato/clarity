local world = require('library.world')
local drawSystem = require('systems.draw')
local playerSystem = require('systems.player')
local uiSystem = require('systems.ui')
local shuffleSystem = require('systems.shuffle')
local playerEntity = require('entities.player')
local jonEntity = require('entities.jon')
local bgEntity = require('entities.bg')

return function()
    local jonX = 200
    local jonY = 200
    local w = world('main')

    w.addSystem(drawSystem())
    w.addSystem(playerSystem())
    w.addSystem(uiSystem())
    w.addSystem(shuffleSystem())

    w.addEntity(bgEntity())
    w.addEntity(playerEntity())

    for index = 1, 5000 do
        w.addEntity(jonEntity(index * 3 + jonX, index * 3 + jonY))
    end

    w.updateSystems()
    w.sortEntities()

    return w
end