local world = require('library.world')
local playerEntity = require('entities.player')
local drawSystem = require('systems.draw')

return function()
    local w = world('main')
    w.addSystem(drawSystem())
    w.addEntity(playerEntity())

    w.updateSystems()
    w.sortEntities()

    return w
end