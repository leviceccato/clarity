local world = require('library.world')
local drawSystem = require('systems.draw')
local playerSystem = require('systems.player')
local playerEntity = require('entities.player')
local bgEntity = require('entities.bg')

return function()
    local w = world('main')
    w.addSystem(drawSystem())
    w.addSystem(playerSystem())
    w.addEntity(bgEntity())
    w.addEntity(playerEntity())

    w.updateSystems()
    w.sortEntities()

    return w
end