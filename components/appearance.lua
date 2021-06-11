local component = require('library.component')

return function(image, quadWidth, quadHeight, sequenceData)
    local imageWidth = image:getWidth()
    local imageHeight = image:getHeight()
    quadWidth = quadWidth or imageWidth
    quadHeight = quadHeight or imageHeight
    sequenceData = sequenceData or {
        default = {
            quadCount = 1,
            shouldLoop = false,
            duration = 0
        }
    }

    local c = component('appearance', {
        image = image,
        quad = 1,
        sequence = 1,
        sequences = {}
    })

    c.data.getImage = function()
        return c.data.sequences[c.data.quad][c.data.sequence]
    end

    local sequence
    for sequenceIndex = 1, #sequenceData do
        sequence = sequenceData[sequenceIndex]
        for frameIndex = 1, sequence.frameCount do
            c.data.sequences[sequenceIndex][frameIndex] = love.graphics.newQuad(
                (frameIndex - 1) * frameWidth,
                (sequenceIndex - 1) * frameHeight,
                frameWidth,
                frameHeight,
                imageWidth,
                imageHeight
            )
        end
    end

    return c
end