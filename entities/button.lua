local entity = require('library.entity')
local appearance = require('components.appearance')
local animation = require('components.animation')
local position = require('components.position')
local text = require('components.text')
local hover = require('components.hover')
local click = require('components.hover')

return function(x, y, content, image, sheet, align, padding, colour)
    local e = entity()

    e.addComponent(appearance(image))
    e.addComponent(animation(sheet))
    e.addComponent(position(x, y))
    e.addComponent(text(content, align, padding, colour))
    e.addComponent(hover())
    e.addComponent(click())

    return e
end