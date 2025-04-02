package iterator

type PlaylistIterator struct {
	playlist *Playlist
	index    int
}

func (p *PlaylistIterator) Next() *Song {
	if p.HasNext() {
		song := p.playlist.Songs[p.index]
		p.index++
		return song
	}
	return nil
}

func (p *PlaylistIterator) Prev() *Song {
	if p.HasPrev() {
		song := p.playlist.Songs[p.index]
		p.index--
		return song
	}
	return nil
}

func (p *PlaylistIterator) HasNext() bool {
	return p.index < len(p.playlist.Songs)
}

func (p *PlaylistIterator) HasPrev() bool {
	return p.index > 0
}

func (p *PlaylistIterator) Reset() {
	p.index = 0
}