local entity = require('library.entity')

local appearance = require('components.appearance')
local position = require('components.position')
local text = require('components.text')
local debug = require('components.debug')

return function()
    local e = entity()

    e.addComponent(appearance(nil, 100, 100))
    e.addComponent(position(10, 10))
    e.addComponent(text())
    e.addComponent(debug())

    return e
end