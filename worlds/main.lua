local world = require('library.world')
local drawSystem = require('systems.draw')
local playerSystem = require('systems.player')
local playerEntity = require('entities.player')

return function()
    local w = world('main')
    w.addSystem(drawSystem())
    w.addSystem(playerSystem())
    w.addEntity(playerEntity())

    w.updateSystems()
    w.sortEntities()

    return w
end